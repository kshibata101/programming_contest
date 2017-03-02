#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N;
  cin >> N;

  int keta = (N-1) / 9;
  int kazu = N % 9;
  if (kazu == 0) {
    kazu = 9;
  }

  for (int i = 0; i <= keta; i++) {
    cout << kazu;
  }
  cout << eol;
  
  return 0;
}
