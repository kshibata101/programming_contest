#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int keta(int val) {
  int k = 1;
  while (val >= 10) {
    k++;
    val /= 10;
  }
  return k;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int a,b;
  cin >> a >> b;
  cout << keta(a) * keta(b) << eol;
  return 0;
}
