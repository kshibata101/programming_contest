#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int L,H,N,A;
  cin >> L >> H >> N;
  for (int i = 0; i < N; i++) {
    cin >> A;
    if (A < L) {
      cout << (L - A) << eol;
    } else if (A <= H) {
      cout << 0 << eol;
    } else {
      cout << -1 << eol;
    }
  }
  
  return 0;
}
