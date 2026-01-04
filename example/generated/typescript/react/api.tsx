// This file is auto-generated. DO NOT EDIT.

import React, { createContext, useContext, useEffect, useRef } from 'react';

const apiContext = createContext({
    baseUrl: '/'
});

type ApiProviderProps = {
    baseUrl: string;
    children: React.ReactNode;
}

export function ApiProvider({ baseUrl, children }: ApiProviderProps) {
    return <apiContext.Provider value={{ baseUrl }}>{children}</apiContext.Provider>;
}

export function useApiBaseUrl() {
    return useContext(apiContext).baseUrl;
}
