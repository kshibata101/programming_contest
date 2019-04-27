#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  vector<int> c = {3,2,1,3,0,2};
  vector<int> v = {1,5,10,50,100,500};
  int a = 620;

  int sum = 0;
  for (int i = c.size() - 1; i >= 0; i--) {
    int num = min(a / v[i], c[i]);
    a -= num * v[i];
    sum += num;
  }

  cout << sum << eol;
  
  return 0;
}
