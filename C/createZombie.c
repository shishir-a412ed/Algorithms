// Create a zombie process in a docker container.
// This program will help reproduce docker PID1 problem.

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/wait.h>

// Adjust sleep time according to your usecase.
int main(){
	pid_t pid, child_pid, grand_child_pid;
	pid=getpid();
	child_pid=fork();
	if (child_pid == -1)
	{
		perror("fork");
		exit(1);
	}

	if (child_pid == 0)
	{
		// Child Process.
		grand_child_pid=fork();
		if (grand_child_pid == -1)
		{
			perror("fork");
			exit(1);
		}

		if (grand_child_pid == 0)
		{
		 // Grand child process.
		}else{
			printf("Grandchild process: %d\n", grand_child_pid);
			int greturnStatus;
			sleep(600);
			waitpid(grand_child_pid, &greturnStatus, 0);
		}
	}else{
		// Parent Process.
		int returnStatus;
		printf("Parent Process: %d\nChild Process: %d\n", pid, child_pid);
		waitpid(child_pid, &returnStatus, 0); 
		sleep(600);
	}
	return 0;
}
