#include <iostream>
#include <math.h>
#include <climits>

using namespace std;

int N;
int* R;

int solve() {
  int mi = R[0];
  int diff = INT_MIN;
  for (int i = 1; i < N; i++) {
    diff = max(diff, R[i] - mi);

    if (R[i] < mi)
      mi = R[i];
  }
  return diff;
}

int main() {
  cin >> N;
  R = new int[N];
  for (int i = 0; i < N; i++) {
    cin >> R[i];
  }

  int res = solve();
  cout << res << endl;

  return 0;
}
