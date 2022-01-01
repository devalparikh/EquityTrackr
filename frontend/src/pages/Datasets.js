import React, { useState, useEffect } from "react";
import PageTitle from "../components/Typography/PageTitle";

import CSVReader2 from "../components/Datasets/CSVReader2";
import CodeEditor from "../components/Code/CodeEditor";
import sampleCode from "../../src/components/Code/sampleCode.txt";
import GenerateCodeButton from "../components/Buttons/GenerateCodeButton";
import DatasetsTable from "../components/Table/DatasetsTable";
import FeaturesTable from "../components/Table/FeaturesTable";

function Datasets() {
  const [pageTable, setPageTable] = useState(1);

  const [dataset, setDataset] = useState([]);
  const [pagedDataset, setPagedDataset] = useState([]);

  const [features, setFeatures] = useState([]);
  const [checkedFeatures, setCheckedFeatures] = useState({});

  const [generatedCode, setGeneratedCode] = useState("");

  // pagination setup
  const resultsPerPage = 10;
  const totalResults = dataset.length - 1;

  // pagination change control
  function onPageChangeTable(p) {
    setPageTable(p);
  }

  // on page change, load new sliced data
  // here you would make another server request for new data
  useEffect(() => {
    setPagedDataset(
      dataset.slice(
        (pageTable - 1) * resultsPerPage,
        pageTable * resultsPerPage
      )
    );
    console.log(
      dataset.slice(
        (pageTable - 1) * resultsPerPage,
        pageTable * resultsPerPage
      )
    );
  }, [pageTable]);

  const updateDataset = (newDataset) => {
    if (newDataset.length < 1) {
      setDataset([]);
      setFeatures([]);
      setCheckedFeatures({});
      setPagedDataset([]);
      return;
    }
    const parsedData = newDataset.map((datapoint) => datapoint.data);
    setDataset(parsedData.slice(1));
    const parsedFeatures = parsedData[0].reduce(
      (o, key, i) => ({ ...o, [key]: true, [i]: true }),
      {}
    );
    setFeatures(parsedData[0]);
    setCheckedFeatures(parsedFeatures);
    setPagedDataset(parsedData.slice(1, 11));
  };

  const updateCheckedFeatures = (feature, index) => {
    console.log({
      ...checkedFeatures,
      [feature]: !checkedFeatures[feature],
      [index]: !checkedFeatures[index],
    });
    setCheckedFeatures({
      ...checkedFeatures,
      [feature]: !checkedFeatures[feature],
      [index]: !checkedFeatures[index],
    });
  };

  const handleGenerateCode = () => {
    fetch(sampleCode)
      .then((r) => r.text())
      .then((text) => {
        setGeneratedCode(text);
      });
  };

  return (
    <>
      <PageTitle>Datasets</PageTitle>

      <CSVReader2 updateDataset={(dataset) => updateDataset(dataset)} />

      {features.length > 0 && (
        <FeaturesTable
          checkedFeatures={checkedFeatures}
          features={features}
          title={"Dataset Features"}
          updateCheckedFeatures={updateCheckedFeatures}
        />
      )}

      {pagedDataset.length > 0 && (
        <>
          <DatasetsTable
            checkedFeatures={checkedFeatures}
            features={features}
            resultsPerPage={resultsPerPage}
            onPageChangeTable={onPageChangeTable}
            pagedDataset={pagedDataset}
            pageTable={pageTable}
            title={"Filtered Dataset"}
            totalResults={totalResults}
          />
          <GenerateCodeButton
            generatedCode={generatedCode}
            handleGenerateCode={handleGenerateCode}
          />
          <CodeEditor code={generatedCode} />
        </>
      )}
    </>
  );
}

export default Datasets;
