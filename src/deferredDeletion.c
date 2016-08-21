// This source code will test whether deferred deletion is supported by underlying kernel.
// This binary should execute successfully when run on fedora or upstream kernels.
// This binary should fail on RHEL kernels.

#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/wait.h>
#include <sys/mount.h>
#include <sched.h>
#include <pthread.h>
#include <semaphore.h>
#include <fcntl.h> 
#include <sys/mman.h>

int main()
{
pid_t parent_pid, child_pid;
int errCode;
int status;
struct stat st = {0};
// Place semaphore in shared memory
sem_t *sem_defer_del = mmap(NULL, sizeof(sem_defer_del), 
      PROT_READ |PROT_WRITE,MAP_SHARED|MAP_ANONYMOUS,
      -1, 0);
// Set pshared flag=1 so that semaphore is shared between parent and child processes.
// pshared=1 is the middle argument in sem_init() call.
sem_init(sem_defer_del, 1, 0);
parent_pid=getpid();
if (stat("/tmp/foo", &st) == -1) {
	mkdir("/tmp/foo", 0700);
}
errCode = mount("/tmp/foo", "/tmp/foo", "", MS_BIND, "");
if (errCode)
{
	printf("Error in mount operation\n");
	perror("mount");
	return 1;
}
child_pid=fork();
if (child_pid == 0)
{
	errCode = unshare(CLONE_NEWNS);
	if (errCode)
	{
		printf("Error unsharing mount namespace\n");
		perror("unshare");
		return 1;
	}
	
	errCode = mount("/", "/", "", MS_PRIVATE|MS_REC, "");
	if (errCode)
	{
        printf("Error in making mount propagation flags PRIVATE\n");
        perror("mount");
	}
	// P() operation of semaphore.
	sem_wait(sem_defer_del);
	sem_destroy(sem_defer_del);
	if (munmap(sem_defer_del, sizeof(sem_defer_del)) < 0) {
      		perror("munmap failed");
      		return 1;
    	}	
}else{
	sleep(5);
	errCode = umount("/tmp/foo");
	if (errCode)
	{
        	printf("Error in unmount operation\n");
        	perror("umount");
        	return 1;
	}
	errCode = rmdir("/tmp/foo");
	if (errCode)
	{
		printf("Error in removing directory in host mount namespace\n");
		perror("rmdir");
		return 1;
	}
	// V() operation of semaphore.
	sem_post(sem_defer_del);
	waitpid(child_pid, &status, 0);
}
return 0;
}
