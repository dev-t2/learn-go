String sayHello({String name = '', required int age}) {
  return 'Hello $name, Are you $age.';
}

int plus(int a, int b) => a + b;

void main(List<String> args) {
  print(sayHello(name: 'Austin', age: 35));
  print(plus(1, 2));
}
