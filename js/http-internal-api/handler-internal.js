exports.handler = async event => {
  return {
    uppercase: event.name.toUpperCase(),
    lowercase: event.name.toLowerCase()
  };
}