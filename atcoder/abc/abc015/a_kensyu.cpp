#include <iostream>

using namespace std;

int main() {
  string a, b;
  cin >> a >> b;

  if (a.length() < b.length()) {
    cout << b << endl;
  } else {
    cout << a << endl;
  }

  return 0;
}
