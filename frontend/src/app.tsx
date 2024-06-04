import { Rover } from './rover';
import './assets/style.css';
import useLocation from './location.hook';

export function App() {
  const { orientation, position } = useLocation();

  return (
    <>
      <div className="container">
        <div>
          <div class="grid"><Rover orientation={orientation} position={position} /></div>
        </div>
        <h2 style="--wails-draggable:drag">Wails + Preact</h2>
      </div>
    </>
  )
}
