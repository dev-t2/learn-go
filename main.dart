String sayHello(String name) {
  return 'Hello $name, Nice to meet you!';
}

int plus(int a, int b) => a + b;

void main(List<String> args) {
  print(sayHello('Austin'));
  print(plus(1, 2));
}
