#include <cstdio>
#include <cstring>
#include <vector>
#include <iostream>
#include <cmath>
#include <algorithm>
using namespace std;
typedef long long ll;

int main(){
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr); 
	cout.tie(nullptr);

	int n;
	cin >> n;
	vector<pair<int, int>> a(n);
	for (auto& p : a) {
		cin >> p.first >> p.second;
	}

	std::sort(a.begin(), a.end());

	int end1 = -1, end2 = -2;
	for (auto&p : a) {
		if (p.first <= end1 && p.first <= end2) {
			cout << "NO";
			return 0;
		}
		else if (p.first > end1) {
			end1 = p.second;
		} else {
			end2 = p.second;
		}
	}
	cout << "YES";
	return 0;
}