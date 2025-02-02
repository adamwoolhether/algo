const std = @import("std");
const stdout = std.io.getStdOut().writer();
const expect = std.testing.expect;

fn bubbleSort(comptime T: type, array: []T) !void {
    var sorted = false;

    while (!sorted) {
        sorted = true;

        var i: u8 = 0;
        while (i < array.len - 1) : (i += 1) {
            if (array[i] > array[i + 1]) {
                const temp = array[i];
                array[i] = array[i + 1];
                array[i + 1] = temp;

                sorted = false;
            }
        }
    }
}

fn selectionSort(comptime T: type, array: []T) !void {
    var i: usize = 0;
    while (i < array.len - 1) : (i += 1) {
        var lowestNumIdx: usize = i;

        var j = i + 1;
        while (j < array.len) : (j += 1) {
            if (array[j] < array[lowestNumIdx])
                lowestNumIdx = j;
        }

        if (lowestNumIdx != i) {
            const temp = array[i];
            array[i] = array[lowestNumIdx];
            array[lowestNumIdx] = temp;
        }
    }
}

fn insertionSort(comptime T: type, array: []T) !void {
    var i: usize = 1;
    while (i < array.len) : (i += 1) {
        const tmp = array[i];
        var j = i;

        while (j > 0 and array[j - 1] > tmp) : (j -= 1) {
            array[j] = array[j - 1];
        }
        array[j] = tmp;
    }
}

// tests

test "bubble sort" {
    const exp = [_]u8{ 10, 15, 25, 35, 45, 55, 65 };
    var input = [_]u8{ 65, 55, 45, 35, 25, 15, 10 };

    try bubbleSort(u8, input[0..]);

    // try stdout.print("{any}\n", .{input});

    try expect(std.mem.eql(u8, exp[0..], input[0..]));
}

test "selection sort" {
    const exp = [_]u8{ 10, 15, 25, 35, 45, 55, 65 };
    var input = [_]u8{ 65, 55, 45, 35, 25, 15, 10 };

    try selectionSort(u8, input[0..]);

    // try stdout.print("{any}\n", .{input});

    try expect(std.mem.eql(u8, exp[0..], input[0..]));
}
test "insertion sort" {
    const exp = [_]u8{ 10, 15, 25, 35, 45, 55, 65 };
    var input = [_]u8{ 65, 55, 45, 35, 25, 15, 10 };

    try insertionSort(u8, input[0..]);

    // try stdout.print("{any}\n", .{input});

    try expect(std.mem.eql(u8, exp[0..], input[0..]));
}
