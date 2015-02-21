#include <iostream>

using namespace std;

int binary_search(int n, int* S, int v) {
  int s = 0, t = n;
  while (s < t) {
    int j = (s + t) / 2;
    if (v == S[j]) {
      return 1;
    } else if (v < S[j]) {
      t = j;
    } else {
      s = j+1;
    }
  }
  return 0;
}



int main() {
  int n,q,v,cnt;
  cin >> n;
  int S[n];
  for (int i = 0; i < n; i++) {
    cin >> S[i];
  }

  cin >> q;
  cnt = 0;
  for (int i = 0; i < q; i++) {
    cin >> v;
    cnt += binary_search(n, S, v);
  }

  cout << cnt << endl;
  return 0;
}
