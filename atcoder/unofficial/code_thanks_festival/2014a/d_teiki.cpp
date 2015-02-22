#include <iostream>
#include <vector>
using namespace std;
     
int main() {
int n, q, a, b, s, t;
 cin >> n >> q;
 for (int i = 0; i < q; i++) {
   cin >> a >> b >> s >> t;
   int ans = 0;
     
   ans += max(min(t, a) - s, 0) + max(t - max(s, b), 0);
   cout << ans * 100 << endl;
 }
     
 return 0;
}
