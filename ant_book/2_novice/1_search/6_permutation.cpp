#include <iostream>

using namespace std;

int n;
bool* used; // array
int* perm;  // array

void permutation(int pos, int n) {
  if (pos == n) {
    // do something for perm which is an array of permutation
    for (int i = 0; i < n; i++) {
      cout << perm[i] << " ";
    }
    cout << endl;
    return;
  }

  // permのpos番目を決定する
  for (int i = 0; i < n; i++) {
    if (used[i])
      continue;

    perm[pos] = i;
    used[i] = true;
    permutation(pos + 1, n);
    used[i] = false; // 終わったらフラグを閉じる
  }
  return;
}

void solve() {
  cin >> n;
  used = new bool[n];
  perm = new int[n];

  // init
  for (int i = 0; i < n; i++) {
    used[i] = false;
    perm[i] = 0;
  }

  permutation(0, n);
  return;
}

int main() {
  solve();
  return 0;
}
