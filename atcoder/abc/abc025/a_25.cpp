#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  string S;
  int N;
  cin >> S >> N;

  int a = (N - 1) / 5;
  int b = (N - 1) % 5;

  cout << S[a] << S[b] << eol;
  
  return 0;
}
