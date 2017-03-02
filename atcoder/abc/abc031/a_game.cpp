#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int a,b;
  cin >> a >> b;

  cout << (max((a+1)*b, a*(b+1))) << eol;
  
  return 0;
}
