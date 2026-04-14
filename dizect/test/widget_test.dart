import 'package:flutter_test/flutter_test.dart';

import 'package:dizect/src/app/app.dart';

void main() {
  testWidgets('shows the professional home screen', (
    WidgetTester tester,
  ) async {
    await tester.pumpWidget(const DistributedStorageApp());

    expect(find.text('Distributed Storage'), findsOneWidget);
    expect(
      find.text('Distributed storage, structured for scale.'),
      findsOneWidget,
    );
  });
}
