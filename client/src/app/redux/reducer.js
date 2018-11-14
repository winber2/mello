const defaults = {

};

export default function appReducer(state = defaults, action) {
  switch(action.type) {
    case 'APP/PLACEHOLDER': {
      return state;
    }
    default:
      return state;
  }
}
