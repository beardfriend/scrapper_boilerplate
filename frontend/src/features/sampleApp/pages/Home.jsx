import React from "react";
import { Tabs, Tab } from "@nextui-org/react";
import TabLayout from "../layouts/TabLayout";
import DB from "../containers/DB";
import Run from "../containers/Run";
function Home() {
  return (
    <>
      <TabLayout>
        <Tabs aria-label="Dynamic tabs" fullWidth className="black">
          <Tab key="run" title="RUN">
            <Run />
          </Tab>
          <Tab key="DB" title="DB">
            <DB />
          </Tab>
        </Tabs>
      </TabLayout>
    </>
  );
}

export default Home;
