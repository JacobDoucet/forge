// This file is auto-generated. DO NOT EDIT.

import { useMemo, useState } from "react";

type ObjWithId = {
  id?: string;
};

export function useFormState<T extends ObjWithId>(initialState: T) {
  const [updates, setUpdates] = useState<T>(() => ({ id: initialState.id } as T));

  const onUpdate = (update: Partial<T>) => {
    setUpdates({
      ...updates,
      ...update,
    });
  };

  const onReset = () => {
    setUpdates({ id: initialState.id } as T);
  };

  const currentState = useMemo(() => {
    return {
      ...initialState,
      ...updates,
    };
  }, [initialState, updates]);

  return { currentState, updates, hasChanges: Object.keys(updates).length > 1, onUpdate, onReset } as const;
}
