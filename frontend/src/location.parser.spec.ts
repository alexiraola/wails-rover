import { describe, it, expect } from "vitest";
import { LocationParser } from "./location.parser";
import { Orientation, Position } from "./location";

describe('Location parser', () => {
  it('should parse different orientations', () => {
    expect(LocationParser.parse('N:0:0').orientation).toBe(Orientation.NORTH);
    expect(LocationParser.parse('E:0:0').orientation).toBe(Orientation.EAST);
    expect(LocationParser.parse('S:0:0').orientation).toBe(Orientation.SOUTH);
    expect(LocationParser.parse('W:0:0').orientation).toBe(Orientation.WEST);
  });

  it('should fail if an invalid orientation is provided', () => {
    expect(() => LocationParser.parse('X:0:0')).toThrow('Invalid orientation: X');
  });

  it('should parse different positions', () => {
    expect(LocationParser.parse('N:1:1').position).toEqual(new Position(1, 1));
    expect(LocationParser.parse('E:10:2').position).toEqual(new Position(10, 2));
    expect(LocationParser.parse('S:11:20').position).toEqual(new Position(11, 20));
    expect(LocationParser.parse('W:15:0').position).toEqual(new Position(15, 0));
  });

  it('should fail if an invalid position is provided', () => {
    expect(() => LocationParser.parse('N:1:a')).toThrow('Invalid position: a');
    expect(() => LocationParser.parse('N:1.5:1')).toThrow('Invalid position: 1.5');
  });
});
