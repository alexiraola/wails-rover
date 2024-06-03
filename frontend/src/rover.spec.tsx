import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/preact';
import { Rover } from './rover';

describe('Rover', () => {
  const getStyles = (location: string) => {
    const { container } = render(<Rover location={location} />);
    const rover = container.querySelector('.rover')!;
    const image = container.querySelector('img')!;
    const roverStyle = getComputedStyle(rover);
    const imgStyle = getComputedStyle(image);

    return [roverStyle, imgStyle];
  }
  it('Should render correctly facing North', () => {
    const [roverStyle, imgStyle] = getStyles('N:0:0');

    expect(roverStyle.transform).toBe('translate(0px, 0px)');
    expect(imgStyle.transform).toBe('rotate(-90deg)');
  });

  it('Should render correctly facing East', () => {
    const [roverStyle, imgStyle] = getStyles('E:0:0');

    expect(roverStyle.transform).toBe('translate(0px, 0px)');
    expect(imgStyle.transform).toBe('rotate(0deg)');
  });

  it('Should render correctly facing South', () => {
    const [roverStyle, imgStyle] = getStyles('S:0:0');

    expect(roverStyle.transform).toBe('translate(0px, 0px)');
    expect(imgStyle.transform).toBe('rotate(90deg)');
  });

  it('Should render correctly facing West', () => {
    const [roverStyle, imgStyle] = getStyles('W:0:0');

    expect(roverStyle.transform).toBe('translate(0px, 0px)');
    expect(imgStyle.transform).toBe('rotate(180deg)');
  });

  it('Should render correctly in the cell [1,1]', () => {
    const [roverStyle, imgStyle] = getStyles('N:1:1');

    expect(roverStyle.transform).toBe('translate(50px, -50px)');
    expect(imgStyle.transform).toBe('rotate(-90deg)');
  });

  it('Should render correctly in the cell [9,9]', () => {
    const [roverStyle, imgStyle] = getStyles('N:9:9');

    expect(roverStyle.transform).toBe('translate(450px, -450px)');
    expect(imgStyle.transform).toBe('rotate(-90deg)');
  });
});
