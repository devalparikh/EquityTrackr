import React from "react";

import {
  Table,
  TableHeader,
  TableCell,
  TableBody,
  TableRow,
  TableFooter,
  TableContainer,
  Pagination,
} from "@windmill/react-ui";

import SectionTitle from "../Typography/SectionTitle";

function DatasetsTable({
  checkedFeatures,
  features,
  resultsPerPage,
  onPageChangeTable,
  pagedDataset,
  pageTable,
  title,
  totalResults,
}) {
  return (
    <>
      <SectionTitle>{title}</SectionTitle>
      <TableContainer className="mb-8">
        <Table>
          <TableHeader>
            <tr>
              <TableCell key={"rowNum"}>#</TableCell>

              {features &&
                features.map(
                  (feature, index) =>
                    checkedFeatures[feature] && (
                      <TableCell key={index}>{feature}</TableCell>
                    )
                )}
            </tr>
          </TableHeader>
          <TableBody>
            {pagedDataset.map((columnArray, colIndex) => (
              <TableRow key={colIndex + "_col"}>
                <TableCell className="bg-gray-50 dark:bg-gray-700">
                  <span
                    key={colIndex + "_rowNum"}
                    className="text-sm text-gray-400"
                  >
                    {colIndex + 1 + (pageTable - 1) * 10}
                  </span>
                </TableCell>
                {columnArray.map(
                  (rowValue, rowIndex) =>
                    checkedFeatures[rowIndex] && (
                      <TableCell>
                        <span
                          key={colIndex + "_" + rowIndex + "_row"}
                          className="text-sm"
                        >
                          {rowValue}
                        </span>
                      </TableCell>
                    )
                )}
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <TableFooter>
          <Pagination
            totalResults={totalResults}
            resultsPerPage={resultsPerPage}
            onChange={onPageChangeTable}
            label="Table navigation"
          />
        </TableFooter>
      </TableContainer>
    </>
  );
}
export default DatasetsTable;
