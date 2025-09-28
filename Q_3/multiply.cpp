#include<iostream>

using namespace std;
int main(){
    int arr[]={7,6,7,6};
    int mul=1;

    for(int i=0; i<4; i++){
        mul = mul * arr[i];
    }
    cout<<mul;
   return 0;
}


