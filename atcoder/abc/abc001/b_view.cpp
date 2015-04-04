#include <iostream>
#include <iomanip>
#define eol '\n';
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int m;
  cin >> m;

  double km = (double)m / 1000.0;

  int ans = 0;
  if (km < 0.1) {
    
  } else if (km <= 5) {
    ans = (int)(km * 10);
  } else if (km <= 30) {
    ans = (int)km + 50;
  } else if (km <= 70) {
    ans = (int)((km - 30) / 5.0) + 80;
  } else {
    ans = 89;
  }
  cout << setfill('0') << setw(2) << ans << eol;
  return 0;
}
