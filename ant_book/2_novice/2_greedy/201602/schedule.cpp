#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n;
  cin >> n;
  vector<int> s(n);
  vector<int> t(n);

  for (int i = 0; i < n; i++) {
    cin >> s[i] >> t[i];
  }

  int now = 0;
  int count = 0;
  
  while (true) {
    int select_end_time = -1;
    
    for (int i = 0; i < n; i++) {
      if (s[i] < now || t[i] < now) {
        continue;
      }
      if (select_end_time == -1 || t[i] < select_end_time) {
        select_end_time = t[i];
      }
    }
    
    if (select_end_time >= 0) {
      count++;
      now = select_end_time + 1;
    } else {
      break;
    }
  }

  cout << count << eol;
  return 0;
}
