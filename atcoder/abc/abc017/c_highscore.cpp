#include <iostream>

using namespace std;

int N, M;
int* l;
int* r;
int* s;

int napsack(int i, bool* gemlist) {
  if (i == N) {
    return 0;
  }

  // don't use
  int res1 = napsack(i + 1, gemlist);
  
  // use 
  bool* gemlist2 = new bool[M];
  int gemCount = 0;
  for (int j = 0; j < M; j++) {
    if (l[i]-1 <= j && j <= r[i]-1) {
      gemlist2[j] = true;
    } else {
      gemlist2[j] = gemlist[j];
    }

    if (gemlist2[j]) {
      gemCount++;
    }
  }

  if (gemCount >= M) { // cannot use
    return res1;
  }

  int res2 = napsack(i + 1, gemlist2);

  return max(res1, res2 + s[i]);
}

int main() {
  cin >> N >> M;

  l = new int[N];
  r = new int[N];
  s = new int[N];
  for (int i = 0; i < N; i++) {
    cin >> l[i] >> r[i] >> s[i];
  }

  bool* gemlist = new bool[M];

  int res = napsack(0, gemlist);
  cout << res << endl;
  return 0;
}
