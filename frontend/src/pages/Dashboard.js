import React, { useState, useEffect } from "react";
import axios from "axios";

import InfoCard from "../components/Cards/InfoCard";
import ChartCard from "../components/Chart/ChartCard";
import { Doughnut, Line } from "react-chartjs-2";
import ChartLegend from "../components/Chart/ChartLegend";
import PageTitle from "../components/Typography/PageTitle";
import { MoneyIcon, ChartsIcon } from "../icons";
import RoundIcon from "../components/RoundIcon";
import response from "../utils/demo/positionsData";

import {
  doughnutOptions,
  lineOptions,
  doughnutLegends,
  lineLegends,
} from "../utils/demo/chartsData";
import { currencyFormat } from "../utils/textFormat";
import PositionsTable from "../components/Table/PositionsTable";

function Dashboard() {
  const [page, setPage] = useState(1);
  const [positions, setPositions] = useState([]);
  const [userData, setUserData] = useState([]);

  useEffect(() => {
    axios
      .get("http://localhost:8080/api/v1/investor?name=Deval")
      .then((response) => {
        setUserData(response.data);
      })
      .catch((err) => {
        console.log(err);
        setUserData({
          balance: 404690,
          email: "mock@gmail.com",
          name: "Mock",
        });
      });
  }, []);

  const resultsPerPage = 4;

  useEffect(() => {
    setPositions(
      response.slice((page - 1) * resultsPerPage, page * resultsPerPage)
    );
  }, [page]);

  return userData.name ? (
    <>
      <PageTitle>Weclome, {userData.name}</PageTitle>

      {/* <!-- Cards --> */}
      <div className="grid gap-6 mb-8 md:grid-cols-2 xl:grid-cols-4">
        <InfoCard
          title="Investment Value"
          value={currencyFormat(userData.balance)}
        >
          <RoundIcon
            icon={MoneyIcon}
            iconColorClass="text-blue-500 dark:text-blue-100"
            bgColorClass="bg-blue-100 dark:bg-blue-500"
            className="mr-4"
          />
        </InfoCard>
        <InfoCard title="Total investments" value={positions.length}>
          <RoundIcon
            icon={ChartsIcon}
            iconColorClass="text-purple-500 dark:text-purple-100"
            bgColorClass="bg-purple-100 dark:bg-purple-500"
            className="mr-4"
          />
        </InfoCard>
        <InfoCard title="Total P/L" value="+$26,360.89 (+17.87%)">
          <RoundIcon
            icon={MoneyIcon}
            iconColorClass="text-green-500 dark:text-green-100"
            bgColorClass="bg-green-100 dark:bg-green-500"
            className="mr-4"
          />
        </InfoCard>
        <InfoCard title="Month P/L" value="+$2,260.89 (+1.68%)">
          <RoundIcon
            icon={MoneyIcon}
            iconColorClass="text-green-500 dark:text-green-100"
            bgColorClass="bg-green-100 dark:bg-green-500"
            className="mr-4"
          />
        </InfoCard>
      </div>

      <div className="grid gap-6 mb-8 md:grid-cols-2">
        <ChartCard title="Value">
          <Line {...lineOptions} />
          <ChartLegend legends={lineLegends} />
        </ChartCard>
        <ChartCard title="Allocation">
          <Doughnut {...doughnutOptions} />
          <ChartLegend legends={doughnutLegends} />
        </ChartCard>
      </div>

      <PositionsTable data={positions} columns={["Name", "Marekt Value"]} />
    </>
  ) : (
    <></>
  );
}

export default Dashboard;
