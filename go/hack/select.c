#include <stdio.h>
#include <sys/time.h>
#include <unistd.h>

int play_select() {
  fd_set rfds;
  FD_ZERO(&rfds);
  FD_SET(STDIN_FILENO, &rfds);

  struct timeval tv;
  tv.tv_sec = 50;
  tv.tv_usec = 0;
  int retval = select(1, &rfds, NULL, NULL, &tv);
  if (retval < 0) {
    printf("select error\n");
    return -1;
  }

  if (retval == 0) {
    printf("time out\n");
    return -1;
  }

  // 需要对所有句柄进行轮询遍历
  if (FD_ISSET(STDIN_FILENO, &rfds)) {
    char buff[255];

    scanf("%s", buff);
    printf("input: %s\n", buff);
  }

  return 0;
}

int main() { play_select(); }
