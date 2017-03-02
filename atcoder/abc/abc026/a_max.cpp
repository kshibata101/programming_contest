#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int A;
  cin >> A;

  int x = A / 2;
  int y = A - x;
  cout << x * y << eol;
  
  return 0;
}
