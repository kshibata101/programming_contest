#include <iostream>
#include <vector>
#include <algorithm>

#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N,A,B;
  cin >> N >> A >> B;

  int X = 0;
  
  for (int i = 0; i < N; i++) {
    string s;
    int d;
    cin >> s >> d;

    int x = min(max(A, d), B);
    if (s == "East") {
      X += x;
    } else if (s == "West"){
      X -= x;
    }
  }

  if (X > 0) {
    cout << "East " << X << eol;
  } else if (X < 0) {
    cout << "West " << -X << eol;
  } else {
    cout << 0 << eol;
  }
  
  return 0;
}
