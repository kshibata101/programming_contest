#include <iostream>
#define eol '\n';
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
  
  int n;
  cin >> n;
  int* a = new int[n+1];
  a[0] = -1;
  for (int i = 1; i <= n; i++) {
    cin >> a[i];
  }
  
  return 0;
}
