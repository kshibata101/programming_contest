#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N, A, B;
  cin >> N >> A >> B;

  bool sente = true;
  if (A == B) {
    if (N % (A+1) == 0) {
      sente = false;
    } else {
      sente = true;
    }
  } else {
    if (A >= N) {
      sente = true;
    } else {
      sente = (A > B);
    }
  }

  if (sente) {
    cout << "Takahashi" << eol;
  } else {
    cout << "Aoki" << eol;
  }
  
  return 0;
}
