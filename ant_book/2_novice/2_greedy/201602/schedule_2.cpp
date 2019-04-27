#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n;
  cin >> n;
  vector< pair<int, int> > time(n);
  for (int i = 0; i < n; i++) {
    cin >> time[i].second >> time[i].first;
  }

  sort(time.begin(), time.end());

  int ans = 0, t = 0;
  for (int i = 0; i < n; i++) {
    if (t < time[i].second) { // start
      ans++;
      t = time[i].first; // end
    }
  }

  cout << ans << eol;
  return 0;
}
