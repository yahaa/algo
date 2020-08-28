#include <stdio.h>

const int N = 1000;
const int BITLEN = 32;
const int BLOCK_SIZE = 100;

int Bucket[1 + N / BLOCK_SIZE] = {0};
int BitMap[1 + BLOCK_SIZE / BITLEN] = {0};

void test() {
  //生成测试数据
  freopen("test.txt", "w", stdout);
  for (int i = 0; i < 1000; ++i) {
    if (i == 127) {
      printf("0\n");
      continue;
    }
    printf("%d\n", i);
  }
  fclose(stdout);

  //读入测试数据
  freopen("test.txt", "r", stdin);
  int Value;
  while (scanf("%d", &Value) != EOF) {
    ++Bucket[Value / BLOCK_SIZE];  //测试数据分段累计
  }
  fclose(stdin);

  freopen("debug.txt", "w", stdout);
  for (int i = 0; i < (1 + N / BLOCK_SIZE); i++) {
    printf("---> %d\n", Bucket[i]);
  }
  fclose(stdout);

  //找出累计计数小于BLOCK_SIZE的
  int Start = -1, i;
  for (i = 0; i < 1 + N / BLOCK_SIZE; ++i) {
    if (Bucket[i] < BLOCK_SIZE) {
      Start = i * BLOCK_SIZE;
      break;
    }
  }
  if (i == 1 + N / BLOCK_SIZE ||
      Bucket[N / BLOCK_SIZE] == 0 && i == N / BLOCK_SIZE)
    return;
  int End = Start + BLOCK_SIZE - 1;

  //在不满足的那段用bitmap来检测
  freopen("test.txt", "r", stdin);
  while (scanf("%d", &Value) != EOF) {
    if (Value >= Start && Value <= End)  // Value必须满足在那段
    {
      int Temp = Value - Start;
      BitMap[Temp / BITLEN] |= (1 << (Temp % BITLEN));
    }
  }
  fclose(stdin);

  //找出不存在的数
  freopen("re.txt", "w", stdout);
  int found = 0;
  for (int i = 0; i < 1 + BLOCK_SIZE / BITLEN; ++i) {
    for (int k = 0; k < BITLEN; ++k) {
      if ((BitMap[i] & (1 << k)) == 0) {
        printf("%d ", i * BITLEN + k + Start);
        found = 1;
        break;
      }
    }
    if (found) break;
  }
  fclose(stdout);
}

int main(int args, char *arg[]) {
  test();
  return 0;
}