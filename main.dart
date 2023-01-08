String sayHello({String name = '', required int age}) {
  return 'Hello $name, Are you $age.';
}

int plus(int a, int b) => a + b;

typedef ListOfInts = List<int>;

ListOfInts reverseList(ListOfInts list) {
  return list.reversed.toList();
}

void main(List<String> args) {
  String result = sayHello(name: 'Austin', age: 35);

  print(result);
}
