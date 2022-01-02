export function currencyFormat(num) {
  let formattedNum;

  formattedNum = "$" + num.toFixed(2).replace(/(\d)(?=(\d{3})+(?!\d))/g, "$1,");
  if (num < 0) {
    formattedNum = "-" + formattedNum.replace("-", "");
  }
  return formattedNum;
}
