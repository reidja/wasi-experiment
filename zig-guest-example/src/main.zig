const std = @import("std");

extern "log" fn log_count(u32, u32) u32;

pub fn main() !void {
    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();
    var res = log_count(2, 3);
    try stdout.print("log_count(2, 3) = {}", .{res});
    try bw.flush();
}
