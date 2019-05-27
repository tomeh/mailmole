export default function (errors) {
  let ret = [];

  Object.values(errors).forEach((element) => {
    ret = ret.concat(element);
  });

  return ret;
}
