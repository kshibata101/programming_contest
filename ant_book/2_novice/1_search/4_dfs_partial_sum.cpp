// 部分和問題
#include <iostream>
#include <cstdio>

int a[4]; 
int n = 4;
int k = 13;

bool dfs(int i, int sum) {
  if (i == n) {
    return sum == k;
  }

  // a[i] を使わない場合
  if (dfs(i+1, sum)) return true;

  // a[i]を使う場合
  if (dfs(i+1, sum + a[i])) return true;

  // 使っても使わなくても作れないので
  return false;
}

int main() {
  a[0] = 1;
  a[1] = 2;
  a[2] = 4;
  a[3] = 7;

  if (dfs(0, 0))
    printf("YES\n");
  else
    printf("NO\n");
  
  return 0;
}
