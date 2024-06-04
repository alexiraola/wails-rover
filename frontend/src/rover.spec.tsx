import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/preact';
import { Rover } from './rover';
import { Orientation, Position } from './location';

describe('Rover', () => {
  const getStyles = (orientation: Orientation, position: Position) => {
    const { container } = render(<Rover orientation={orientation} position={position} />);
    const rover = container.querySelector('.rover')!;
    const image = container.querySelector('img')!;
    const roverStyle = getComputedStyle(rover);
    const imgStyle = getComputedStyle(image);

    return [roverStyle, imgStyle];
  }
  it('Should render correctly facing North', () => {
    const [roverStyle, imgStyle] = getStyles(Orientation.NORTH, new Position(0, 0));

    expect(roverStyle.transform).toBe('translate(0px, 0px)');
    expect(imgStyle.transform).toBe('rotate(-90deg)');
  });

  it('Should render correctly facing East', () => {
    const [roverStyle, imgStyle] = getStyles(Orientation.EAST, new Position(0, 0));

    expect(roverStyle.transform).toBe('translate(0px, 0px)');
    expect(imgStyle.transform).toBe('rotate(0deg)');
  });

  it('Should render correctly facing South', () => {
    const [roverStyle, imgStyle] = getStyles(Orientation.SOUTH, new Position(0, 0));

    expect(roverStyle.transform).toBe('translate(0px, 0px)');
    expect(imgStyle.transform).toBe('rotate(90deg)');
  });

  it('Should render correctly facing West', () => {
    const [roverStyle, imgStyle] = getStyles(Orientation.WEST, new Position(0, 0));

    expect(roverStyle.transform).toBe('translate(0px, 0px)');
    expect(imgStyle.transform).toBe('rotate(180deg)');
  });

  it('Should render correctly in the cell [1,1]', () => {
    const [roverStyle, imgStyle] = getStyles(Orientation.NORTH, new Position(1, 1));

    expect(roverStyle.transform).toBe('translate(50px, -50px)');
    expect(imgStyle.transform).toBe('rotate(-90deg)');
  });

  it('Should render correctly in the cell [9,9]', () => {
    const [roverStyle, imgStyle] = getStyles(Orientation.NORTH, new Position(9, 9));

    expect(roverStyle.transform).toBe('translate(450px, -450px)');
    expect(imgStyle.transform).toBe('rotate(-90deg)');
  });
});
