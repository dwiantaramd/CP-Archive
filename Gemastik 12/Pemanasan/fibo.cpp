#include <bits/stdc++.h>

using namespace std;

int main() {
    int value;
    int x[10];
    int i = 0;
    while(cin >> value) {
        x[i] = value;
        i++;
    }
    bool check = true;
    for(int idx = i-1; idx >= 2; idx--) {
        if(x[idx-1] + x[idx-2] != x[idx]) {
            check = false;
        }
    }
    if(check) {
        cout << "BENAR";
    } else {
        cout << "SALAH";
    }

}
