import { Rover } from './rover';
import './assets/style.css';
import useLocation from './location.hook';

export function App() {
  const { orientation, position, worldSize } = useLocation();

  const width = worldSize[0] * 50 + 1;
  const height = worldSize[1] * 50 + 1;

  return (
    <>
      <div className="container" style="--wails-draggable:drag">
        <div>
          <div class="grid" style={{ width, height }}><Rover orientation={orientation} position={position} /></div>
        </div>
      </div>
    </>
  )
}
