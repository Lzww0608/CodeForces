#include <iostream>
#include <unordered_map>
#include <vector>
#include <utility>
#include <map>
using ll = long long;
using namespace std;


int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int t;
	cin >> t;
	while (t--) {
		ll n, x;
		cin >> n >> x;
		ll res = n;
		while (res > x) {
			n += n & -n;
			res &= n;
		}
		if (res == x) {
			cout << n << endl;
		} else {
			cout << -1 << endl;
		}
	}
	return 0;

}