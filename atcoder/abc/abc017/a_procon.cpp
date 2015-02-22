#include <iostream>

using namespace std;

int main() {
  int s1,e1,s2,e2,s3,e3;
  cin >> s1 >> e1 >> s2 >> e2 >> s3 >> e3;
  int res = s1 * e1 / 10 + s2 * e2 / 10 + s3 * e3 / 10;
  cout << res << endl;
  return 0;
}
