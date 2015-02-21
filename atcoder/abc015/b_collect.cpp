#include <iostream>
#include <math.h>

using namespace std;

int main() {
  int n, a;
  cin >> n;

  int cnt = 0;
  int sum = 0;

  for (int i = 0; i < n; ++i) {
    cin >> a;
    
    if (a > 0) {
      ++cnt;
      sum += a;
    }

  }

  int ans = ceil((double)sum / cnt);

  cout << ans << endl;
  
  return 0;
}
