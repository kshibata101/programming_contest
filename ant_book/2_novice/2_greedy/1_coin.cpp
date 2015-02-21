#include <iostream>

using namespace std;

int n = 6;
int A;
int* c;
int v[] = {1, 5, 10, 50, 100, 500};

int solve() {
  int res = 0;
  int remain = A;
  for (int i = n-1; i >=0; i--) {
    int count = min(c[i], remain / v[i]);
    res += count;
    remain -= count * v[i];
  }
  return res;
}

int main() {
  c = new int[n];

  // .inファイルからデータ取得
  for (int i = 0; i < n; i++) {
    cin >> c[i];
  }
  cin >> A;

  int res = solve();
  cout << res << endl;

  return 0;
}
