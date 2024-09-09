#include <iostream>
#include <unordered_map>
#include <vector>
#include <utility>
#include <map>
using ll = long long;
using namespace std;

struct Point {
	ll x, y;
	Point(ll a, ll b): x(a), y(b) {}
	bool operator==(const Point& other) const {
		return other.x == x && other.y == y;
	}
};

namespace std {
	template <>
	struct hash<Point> {
		ll operator()(const Point& p) const {
			return (hash<ll>()(p.x) << 9) | hash<ll>()(p.y);
		}
	};
}

int main() {
	int n;
	cin >> n;
	int x, y;

	unordered_map<int, int> m1;
	unordered_map<int, int> m2;
	//unordered_map<Point, ll> point;
	map<pair<int, int>, int> mp;
	ll ans = 0;

	for (int i = 0; i < n; ++i) {
		cin >> x >> y;
		//ans += m1[x]++ + m2[y]++ - point[Point(x, y)]++;
		ans += m1[x]++ + m2[y]++ - mp[{x, y}]++;
	}


	cout << ans << endl;
	return 0;

}
