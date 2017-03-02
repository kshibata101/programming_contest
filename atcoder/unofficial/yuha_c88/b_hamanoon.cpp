#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n;
  cin >> n;

  if (n % 4 == 0) {
    cout << "GO" << eol;
  } else {
    cout << "SEN" << eol;
  }
  
  return 0;
}
