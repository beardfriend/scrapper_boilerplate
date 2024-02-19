import React from "react";
import { RecoilRoot } from "recoil";
import { QueryClient, QueryClientProvider } from "react-query";
import { HashRouter, Route, Routes } from "react-router-dom";
import MetaLayout from "./shared/layout/MetaLayout";
import SampleApp from "./features/sampleApp/pages/Home";
import { NextUIProvider } from "@nextui-org/react";

function App() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        refetchOnWindowFocus: false,
        refetchOnReconnect: false,
        retry: false,
        staleTime: Infinity,
      },
    },
  });
  return (
    <NextUIProvider>
      <RecoilRoot>
        <QueryClientProvider client={queryClient}>
          <MetaLayout>
            <HashRouter basename={"/"}>
              <Routes>
                <Route path="/" element={<SampleApp />} />
              </Routes>
            </HashRouter>
          </MetaLayout>
        </QueryClientProvider>
      </RecoilRoot>
    </NextUIProvider>
  );
}

export default App;
