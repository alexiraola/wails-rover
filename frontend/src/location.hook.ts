import { Events } from "@wailsio/runtime";
import { useEffect, useState } from "preact/hooks";

export default function useLocation() {
  const [location, setLocation] = useState('N:0:0');

  useEffect(() => {
    Events.On('location', (message: any) => {
      setLocation(message.data);
    });
  }, []);

  return { location }
}
