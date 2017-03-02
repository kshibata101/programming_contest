#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N;
  cin >> N;
  
  vector<int> K(N-1);
  for (int i = 0; i < N-1; i++) {
    cin >> K[i];
  }

  cout << K[0];
  
  for (int i = 1; i < N-1; i++) {
    cout << " "  << min(K[i], K[i-1]);
  }
  cout << " " << K[N-2] << eol;
  
  return 0;
}
