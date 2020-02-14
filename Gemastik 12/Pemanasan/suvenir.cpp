#include <bits/stdc++.h>

using namespace std;

int main() {
    int value,i,temp;
    int x[10];
    int y[10];
    i = 0;
    while(cin >> value) {
        x[i] = value;
        y[i] = value;
        i++;
    }
    for(int idx = 0; idx < i-1; idx++){
        temp = x[idx] % x[idx+1];
        while(temp != 0)
        {
            x[idx] = x[idx+1];
            x[idx+1] = temp;
            temp = x[idx] % x[idx+1];
        }
    }
    cout << x[i-1] << endl;
    for(int idx = 0; idx < i; idx++) {
        cout << y[idx]/x[i-1] << " ";
    }
}
