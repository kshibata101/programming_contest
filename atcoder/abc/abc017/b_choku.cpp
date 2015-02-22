#include <iostream>

using namespace std;

int main() {
  string x;
  cin >> x;

  bool res = true;
  int i = 0;
  while(i < x.size()) {
    char c = x[i];
    if (c == 'c') {
      i++;
      if (i >= x.size()) {
        res = false;
        break;
      }
      c = x[i];
      if (c != 'h') {
        res = false;
        break;
      }
    } else if (c == 'o' || c == 'k' || c == 'u') {
    } else {
      res = false;
      break;
    }
    i++;
  }

  if (res) {
    cout << "YES" << endl;
  } else {
    cout << "NO" << endl;
  }

  return 0;
}
