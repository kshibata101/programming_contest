#include <iostream>

using namespace std;

int n, q;
int* S;
int* T;

int solve() {
  int count = 0;
  for (int i = 0; i < q; i++) {
    int t = T[i];
    
    int left = 0;
    int right = n;
    while (left < right) {
      int mid = (left + right) / 2;
      if (S[mid] == t) {
        count++;
        break;
      } else if (S[mid] < t) {
        left = mid + 1; // midはもう調べたので含めない
      } else {
        right = mid; // rightには含めない
      }
    }
  }
  return count;
}

int main() {
  cin >> n;
  S = new int[n];
  for (int i = 0; i < n; i++) {
    cin >> S[i];
  }
  
  cin >> q;
  T = new int[q];
  for (int i = 0; i < q; i++) {
    cin >> T[i];
  }

  int res = solve();
  cout << res << endl;

  return 0;
}
